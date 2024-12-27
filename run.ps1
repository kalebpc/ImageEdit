Add-Type -AssemblyName System.Windows.Forms, PresentationFramework

function Initialize-XamlWindow {
    [CmdletBinding()]
    [OutputType([psobject])]
    param (
        [string]$Path,
        [string]$VariablePrefix
    )
    [string]$xamlContent = Get-Content $Path
    [string]$xamlContent = $xamlContent -replace 'mc:Ignorable="d"', '' -replace 'x:N', 'N' -replace 'x:Class=".*?"', ''
    [xml]$xmlNodes = $xamlContent

    $reader = New-Object System.Xml.XmlNodeReader $xmlNodes

    $window = [Windows.Markup.XamlReader]::Load($reader)

    $variables = @{}

    $xmlNodes.SelectNodes("//*[@Name]") | ForEach-Object { Set-Variable -Name (-join ("xaml", $VariablePrefix, $_.Name)) -Value $window.FindName($_.Name) }

    Get-Variable xaml* | ForEach-Object { $variables[$_.Name] = $_.Value }

    return @{window = [System.Windows.Window]$window; variables = [hashtable]$variables}
}

$MainWindow = Initialize-XamlWindow $(Join-Path -Path $(Split-Path -Path "./MainWindow.xaml" -Parent -Resolve) -ChildPath "MainWindow.xaml") "Main"

$DataTable = @{
    imageFileName = " "
    inputImageFullPath = " "
    outputImageFullPath = " "
    outputImageTempPath = "$HOME\AppData\Local\Temp\imageedit\imageedit$("{0:yyyyMMdd}" -f (Get-Date)).png"
    chosenTransformation = " "
    pixels = 0
}

If (Test-Path -Path $(Split-Path -Path $DataTable.outputImageTempPath -Parent)) {
    Remove-Item -Path $(-join ($(Split-Path -Path $DataTable.outputImageTempPath -Parent), "\*")) -Recurse
} Else {
    New-Item -ItemType Directory -Path $(Split-Path -Path $DataTable.outputImageTempPath -Parent)
}

$MainWindow.variables.xamlMainMenuPanel.Add_MouseDown({
    $MainWindow.window.SizeToContent = "Manual"
    $MainWindow.window.WindowState = "Normal"
    $MainWindow.window.DragMove()
})

$MainWindow.variables.xamlMainMenuPanel.Add_MouseUp({
    $MainWindow.window.WindowState = "Normal"
    $MainWindow.window.SizeToContent = "WidthAndHeight"
})

$MainWindow.variables.xamlMainMenuWindowClose.Add_Click({
    $MainWindow.window.DialogResult = [System.Windows.Forms.DialogResult]::Cancel
})

$MainWindow.variables.xamlMainWindowControlClose.Add_Click({
    $MainWindow.window.DialogResult = [System.Windows.Forms.DialogResult]::Cancel
})

$MainWindow.variables.xamlMainAboutLink.Add_Click({
    [System.Diagnostics.Process]::Start($MainWindow.variables.xamlMainAboutLink.Header.substring(17))
})

$MainWindow.variables.xamlMainBrowseButton.Add_Click({
    $openFileDialog = New-Object System.Windows.Forms.OpenFileDialog
    $openFileDialog.InitialDirectory = "%USERPROFILE%\Pictures"
    $openFileDialog.Filter = "png files (*.png)|*.png"
    $openFileDialog.FilterIndex = 2
    $openFileDialog.CheckPathExists
    If ($openFileDialog.ShowDialog() -eq [System.Windows.Forms.DialogResult]::OK) {
        $MainWindow.variables.xamlMainBrowseTextBox.Text = $openFileDialog.FileName
        $DataTable.inputImageFullPath = $openFileDialog.FileName
        $DataTable.imageFileName = Split-Path $openFileDialog.FileName -Leaf
        $bitmapImage = New-Object System.Windows.Media.Imaging.BitmapImage
        $bitmapImage.BeginInit()
        $bitmapImage.DecodePixelWidth = 250
        $bitmapImage.UriSource = New-Object System.Uri($DataTable.inputImageFullPath)
        $bitmapImage.EndInit()
        $MainWindow.variables.xamlMainInputImage.Source = $bitmapImage
        Copy-Item -Path $DataTable.inputImageFullPath -Destination $DataTable.outputImageTempPath
    }
})

$functionList = @("Flipx", "Flipy", "Rotate", "Roundrobinx", "Roundrobiny", "Roundrobinrows", "Roundrobincolumns", "Pixelate", "Rgbfilter")
ForEach ($function in $functionList) {
    $MainWindow.variables.xamlMainTransformationComboBox.Items.Add($function)
}
$MainWindow.variables.xamlMainTransformationComboBox.Add_DropDownClosed({
    $DataTable.chosenTransformation = $MainWindow.variables.xamlMainTransformationComboBox.SelectedItem
    If ($DataTable.chosenTransformation -in @("Flipx", "Flipy", "Rotate")) {
        $MainWindow.variables.xamlMainPixelSlider.Visibility = "Collapsed"
        $MainWindow.variables.xamlMainPixelLabel.Visibility = "Collapsed"
        $MainWindow.variables.xamlMainColorLabel.Visibility = "Collapsed"
        $MainWindow.variables.xamlMainColorComboBox.Visibility = "Collapsed"
        $MainWindow.variables.xamlMainPixelDisplayLabel.Visibility = "Collapsed"
    } ElseIf ($DataTable.chosenTransformation -in @("Roundrobinx", "Roundrobiny", "Roundrobinrows", "Roundrobincolumns", "Pixelate")) {
        $MainWindow.variables.xamlMainColorLabel.Visibility = "Collapsed"
        $MainWindow.variables.xamlMainColorComboBox.Visibility = "Collapsed"
        $MainWindow.variables.xamlMainPixelSlider.Visibility = "Visible"
        $MainWindow.variables.xamlMainPixelLabel.Visibility = "Visible"
        $MainWindow.variables.xamlMainPixelDisplayLabel.Visibility = "Visible"
    } Else {
        $MainWindow.variables.xamlMainPixelSlider.Visibility = "Collapsed"
        $MainWindow.variables.xamlMainPixelLabel.Visibility = "Collapsed"
        $MainWindow.variables.xamlMainPixelDisplayLabel.Visibility = "Collapsed"
        $MainWindow.variables.xamlMainColorLabel.Visibility = "Visible"
        $MainWindow.variables.xamlMainColorComboBox.Visibility = "Visible"
    }
    $MainWindow.variables.xamlMainPreviewButton.Visibility = "Visible"
})

$MainWindow.variables.xamlMainPixelSlider.Maximum = 100
$MainWindow.variables.xamlMainPixelSlider.LargeChange = 10
$MainWindow.variables.xamlMainPixelSlider.SmallChange = 1
$MainWindow.variables.xamlMainPixelSlider.Add_ValueChanged({
    $DataTable.pixels = [Math]::round($MainWindow.variables.xamlMainPixelSlider.Value)
    $MainWindow.variables.xamlMainPixelDisplayLabel.Content = $DataTable.pixels
})

ForEach ($color in @("Red", "Green", "Blue")) {
    $MainWindow.variables.xamlMainColorComboBox.Items.Add($color)
}
$MainWindow.variables.xamlMainColorComboBox.Add_DropDownClosed({
    If ($MainWindow.variables.xamlMainColorComboBox.SelectedItem -match "Red") {
        $DataTable.pixels = 1
    } ElseIf ($MainWindow.variables.xamlMainColorComboBox.SelectedItem -match "Green") {
        $DataTable.pixels = 2
    } Else {
        $DataTable.pixels = 3
    }
})

$MainWindow.variables.xamlMainPreviewButton.Add_Click({
    If (Test-Path -Path $DataTable.outputImageTempPath) {
        ./ImageEdit-amd64 $(-join ("infile", "=", '"', $($DataTable.inputImageFullPath), '"')) $(-join ("outfile", "=", '"', $($DataTable.outputImageTempPath), '"')) $(("function", $DataTable.chosenTransformation) -join "=") $(("pixels", $DataTable.pixels) -join "=")
        Display-ImagePreview
    }
})

$MainWindow.variables.xamlMainSaveButton.Add_Click({
    If ($DataTable.inputImageFullPath -ne " ") {
        $saveFileDialog = New-Object System.Windows.Forms.SaveFileDialog
        $saveFileDialog.InitialDirectory = "%USERPROFILE%\Pictures"
        $saveFileDialog.Filter = "png files (*.png)|*.png"
        $saveFileDialog.FilterIndex = 2
        $saveFileDialog.OverwritePrompt
        If ($DataTable.chosenTransformation -ne " ") {
            If ($DataTable.chosenTransformation -in @("Roundrobinx", "Roundrobiny", "Roundrobinrows", "Roundrobincolumns", "Pixelate")) {
                $saveFileDialog.FileName = -join ( $([System.IO.Path]::GetFileNameWithoutExtension($DataTable.inputImageFullPath)), "-", $DataTable.chosenTransformation, "-", $DataTable.pixels )
            } ElseIf ($DataTable.chosenTransformation -eq "Rgbfilter") {
                $color = ""
                If ($DataTable.pixels -eq 1) {
                    $color = "Red"
                } ElseIf ($DataTable.pixels -eq 2) {
                    $color = "Green"
                } Else {
                    $color = "blue"
                }
                $saveFileDialog.FileName = -join ( $([System.IO.Path]::GetFileNameWithoutExtension($DataTable.inputImageFullPath)), "-", $DataTable.chosenTransformation, "-", $color )
            }            
        } Else {
            $saveFileDialog.FileName = -join ($([System.IO.Path]::GetFileNameWithoutExtension($DataTable.inputImageFullPath)), "(copy)")
        }
        If ($saveFileDialog.ShowDialog() -eq [System.Windows.Forms.DialogResult]::OK) {
            Copy-Item -Path $DataTable.outputImageTempPath -Destination $saveFileDialog.FileName
        }
    }
})

function Display-ImagePreview {
    $PreviewWindow = Initialize-XamlWindow $(Join-Path -Path $(Split-Path -Path "./PreviewWindow.xaml" -Parent -Resolve) -ChildPath "PreviewWindow.xaml") "Preview"

    $bitmapImage = New-Object System.Windows.Media.Imaging.BitmapImage
    $bitmapImage.BeginInit()
    $bitmapImage.DecodePixelWidth = 500
    $bitmapImage.UriSource = New-Object System.Uri($DataTable.outputImageTempPath)
    $bitmapImage.EndInit()
    $PreviewWindow.variables.xamlPreviewOutputImage.Source = $bitmapImage

    $PreviewWindow.variables.xamlPreviewMenuPanel.Add_MouseDown({
        $PreviewWindow.window.SizeToContent = "Manual"
        $PreviewWindow.window.WindowState = "Normal"
        $PreviewWindow.window.DragMove()
    })

    $PreviewWindow.variables.xamlPreviewMenuPanel.Add_MouseUp({
        $PreviewWindow.window.WindowState = "Normal"
        $PreviewWindow.window.SizeToContent = "WidthAndHeight"
    })

    $PreviewWindow.variables.xamlPreviewWindowControlClose.Add_Click({
        $PreviewWindow.window.DialogResult = [System.Windows.Forms.DialogResult]::Cancel
    })

    $PreviewWindow.variables.xamlPreviewAboutLink.Add_Click({
        [System.Diagnostics.Process]::Start($PreviewWindow.variables.xamlPreviewAboutLink.Header.substring(17))
    })

    $PreviewWindow.window.ShowDialog() | Out-Null
    
    return $Null
}

$MainWindow.window.ShowDialog() | Out-Null
