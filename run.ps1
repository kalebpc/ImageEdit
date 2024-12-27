Add-Type -AssemblyName System.Windows.Forms, PresentationFramework

function Initialize-XamlWindow {
    [CmdletBinding()]
    [OutputType([psobject])]
    param ([string]$Path)
    [string]$xamlContent = Get-Content $Path
    [string]$xamlContent = $xamlContent -replace 'mc:Ignorable="d"', '' -replace 'x:N', 'N' -replace 'x:Class=".*?"', ''
    [xml]$xmlNodes = $xamlContent

    $reader = New-Object System.Xml.XmlNodeReader $xmlNodes

    $window = [Windows.Markup.XamlReader]::Load($reader)

    $variables = @{}

    $xmlNodes.SelectNodes("//*[@Name]") | ForEach-Object { Set-Variable -Name ("xaml", $_.Name -join "") -Value $window.FindName($_.Name) }

    Get-Variable xaml* | ForEach-Object { $variables[$_.Name] = $_.Value }

    return @{window = [System.Windows.Window]$window; variables = [hashtable]$variables}
}

$Window = Initialize-XamlWindow $(Join-Path -Path $(Split-Path -Path "./MainWindow.xaml" -Parent -Resolve) -ChildPath "MainWindow.xaml")

$DataTable = @{
    imageFileName = " "
    inputImageFullPath = " "
    outputImageFullPath = " "
    outputImageTempPath = "$HOME\AppData\Local\Temp\imageedit\imageedit$("{0:yyyyMMdd}" -f (Get-Date)).png"
    chosenTransformation = " "
    pixels = 0
}

If (Test-Path -LiteralPath $(Split-Path -Path $DataTable.outputImageTempPath -Parent)) {
    Remove-Item -Path $(-join ($(Split-Path -Path $DataTable.outputImageTempPath -Parent), "\*")) -Recurse
} Else {
    New-Item -ItemType Directory -Path $(Split-Path -Path $DataTable.outputImageTempPath -Parent)
}

$Window.variables.xamlMenuPanel.Add_MouseDown({
    $Window.window.SizeToContent = "Manual"
    $Window.window.WindowState = "Normal"
    $Window.window.DragMove()
})

$Window.variables.xamlMenuPanel.Add_MouseUp({
    $Window.window.WindowState = "Normal"
    $Window.window.SizeToContent = "WidthAndHeight"
})

$Window.variables.xamlMenuWindowClose.Add_Click({
    $Window.window.DialogResult = [System.Windows.Forms.DialogResult]::Cancel
})

$Window.variables.xamlWindowControlClose.Add_Click({
    $Window.window.DialogResult = [System.Windows.Forms.DialogResult]::Cancel
})

$Window.variables.xamlAboutLink.Add_Click({
    [System.Diagnostics.Process]::Start($Window.variables.xamlAboutLink.Header.substring(17))
})

$Window.variables.xamlBrowseButton.Add_Click({
    $openFileDialog = New-Object System.Windows.Forms.OpenFileDialog
    $openFileDialog.InitialDirectory = "%USERPROFILE%\Pictures"
    $openFileDialog.Filter = "png files (*.png)|*.png"
    $openFileDialog.FilterIndex = 2
    $openFileDialog.CheckPathExists
    If ($openFileDialog.ShowDialog() -eq [System.Windows.Forms.DialogResult]::OK) {
        $Window.variables.xamlBrowseTextBox.Text = $openFileDialog.FileName
        $DataTable.inputImageFullPath = $openFileDialog.FileName
        $DataTable.imageFileName = Split-Path $openFileDialog.FileName -Leaf
        $bitmapImage = New-Object System.Windows.Media.Imaging.BitmapImage
        $bitmapImage.BeginInit()
        $bitmapImage.DecodePixelWidth = 200
        $bitmapImage.UriSource = New-Object System.Uri($DataTable.inputImageFullPath)
        $bitmapImage.EndInit()
        $Window.variables.xamlInputImage.Source = $bitmapImage
        $bitmapImage.Dispose()
        Copy-Item -Path $DataTable.inputImageFullPath -Destination $DataTable.outputImageTempPath
    }
})

$functionList = @("Flipx", "Flipy", "Rotate", "Roundrobinx", "Roundrobiny", "Roundrobinrows", "Roundrobincolumns", "Pixelate", "Rgbfilter")
ForEach ($function in $functionList) {
    $Window.variables.xamlTransformationComboBox.Items.Add($function)
}
$Window.variables.xamlTransformationComboBox.Add_DropDownClosed({
    $DataTable.chosenTransformation = $Window.variables.xamlTransformationComboBox.SelectedItem
    If ($DataTable.chosenTransformation -in @("Flipx", "Flipy", "Rotate")) {
        $Window.variables.xamlPixelSlider.Visibility = "Collapsed"
        $Window.variables.xamlPixelLabel.Visibility = "Collapsed"
        $Window.variables.xamlColorLabel.Visibility = "Collapsed"
        $Window.variables.xamlColorComboBox.Visibility = "Collapsed"
        $Window.variables.xamlPixelDisplayLabel.Visibility = "Collapsed"
    } ElseIf ($DataTable.chosenTransformation -in @("Roundrobinx", "Roundrobiny", "Roundrobinrows", "Roundrobincolumns", "Pixelate")) {
        $Window.variables.xamlColorLabel.Visibility = "Collapsed"
        $Window.variables.xamlColorComboBox.Visibility = "Collapsed"
        $Window.variables.xamlPixelSlider.Visibility = "Visible"
        $Window.variables.xamlPixelLabel.Visibility = "Visible"
        $Window.variables.xamlPixelDisplayLabel.Visibility = "Visible"
    } Else {
        $Window.variables.xamlPixelSlider.Visibility = "Collapsed"
        $Window.variables.xamlPixelLabel.Visibility = "Collapsed"
        $Window.variables.xamlPixelDisplayLabel.Visibility = "Collapsed"
        $Window.variables.xamlColorLabel.Visibility = "Visible"
        $Window.variables.xamlColorComboBox.Visibility = "Visible"
    }
    If ($Window.variables.xamlPreviewButton.Visibility -eq "Hidden") {
        $Window.variables.xamlPreviewButton.Visibility = "Visible"
    }    
})

$Window.variables.xamlPixelSlider.Maximum = 100
$Window.variables.xamlPixelSlider.LargeChange = 10
$Window.variables.xamlPixelSlider.SmallChange = 1
$Window.variables.xamlPixelSlider.Add_ValueChanged({
    $DataTable.pixels = [Math]::round($Window.variables.xamlPixelSlider.Value)
    $Window.variables.xamlPixelDisplayLabel.Content = $DataTable.pixels
})

ForEach ($color in @("Red", "Green", "Blue")) {
    $Window.variables.xamlColorComboBox.Items.Add($color)
}
$Window.variables.xamlColorComboBox.Add_DropDownClosed({
    If ($Window.variables.xamlColorComboBox.SelectedItem -match "Red") {
        $DataTable.pixels = 1
    } ElseIf ($Window.variables.xamlColorComboBox.SelectedItem -match "Green") {
        $DataTable.pixels = 2
    } Else {
        $DataTable.pixels = 3
    }
})

$Window.variables.xamlPreviewButton.Add_Click({
    If (Test-Path -Path $DataTable.outputImageTempPath) {
        ./ImageEdit-amd64 $(-join ("infile", "=", '"', $($DataTable.inputImageFullPath), '"')) $(-join ("outfile", "=", '"', $($DataTable.outputImageTempPath), '"')) $(("function", $DataTable.chosenTransformation) -join "=") $(("pixels", $DataTable.pixels) -join "=")
        $bitmapImage = New-Object System.Windows.Media.Imaging.BitmapImage
        $bitmapImage.BeginInit()
        $bitmapImage.DecodePixelWidth = 200
        $bitmapImage.UriSource = New-Object System.Uri($DataTable.outputImageTempPath)
        $bitmapImage.EndInit()
        $Window.variables.xamlOutputImage.Source = $bitmapImage
        $Window.variables.xamlPreviewButton.Visibility = "Collapsed"
    }
})

$Window.variables.xamlSaveButton.Add_Click({
    If ($DataTable.inputImageFullPath -ne " ") {
        $saveFileDialog = New-Object System.Windows.Forms.SaveFileDialog
        $saveFileDialog.InitialDirectory = "%USERPROFILE%\Pictures"
        $saveFileDialog.Filter = "png files (*.png)|*.png"
        $saveFileDialog.FilterIndex = 2
        $saveFileDialog.OverwritePrompt
        If ($DataTable.chosenTransformation -ne " ") {
            $saveFileDialog.FileName = -join ( $([System.IO.Path]::GetFileNameWithoutExtension($DataTable.inputImageFullPath)), "-", $DataTable.chosenTransformation )
        } Else {
            $saveFileDialog.FileName = -join ($([System.IO.Path]::GetFileNameWithoutExtension($DataTable.inputImageFullPath)), "(copy)")
        }
        If ($saveFileDialog.ShowDialog() -eq [System.Windows.Forms.DialogResult]::OK) {
            Copy-Item -Path $DataTable.outputImageTempPath -Destination $saveFileDialog.FileName
        }
    }
})

$Window.window.ShowDialog() | Out-Null
