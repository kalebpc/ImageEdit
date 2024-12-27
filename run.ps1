Add-Type -AssemblyName System.Windows.Forms, PresentationFramework

[string]$xamlContent = Get-Content $(Join-Path -Path $(Split-Path -Path "./MainWindow.xaml" -Parent -Resolve) -ChildPath "MainWindow.xaml")
[string]$xamlContent = $xamlContent -replace 'mc:Ignorable="d"', '' -replace 'x:N', 'N' -replace 'x:Class=".*?"', ''
[xml]$xmlNodes = $xamlContent

$reader = New-Object System.Xml.XmlNodeReader $xmlNodes

$window = [Windows.Markup.XamlReader]::Load($reader)

$variables = @{}

$xmlNodes.SelectNodes("//*[@Name]") | ForEach-Object { Set-Variable -Name (-join ("xaml", $_.Name)) -Value $window.FindName($_.Name) }

Get-Variable xaml* | ForEach-Object { $variables[$_.Name] = $_.Value }

$MainWindow = @{window = [System.Windows.Window]$window; variables = [hashtable]$variables}

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

$MainWindow.variables.xamlMenuPanel.Add_MouseDown({
    $MainWindow.window.SizeToContent = "Manual"
    $MainWindow.window.WindowState = "Normal"
    $MainWindow.window.DragMove()
})

$MainWindow.variables.xamlMenuPanel.Add_MouseUp({
    $MainWindow.window.WindowState = "Normal"
    $MainWindow.window.SizeToContent = "WidthAndHeight"
})

$MainWindow.variables.xamlMenuWindowClose.Add_Click({
    $MainWindow.window.DialogResult = [System.Windows.Forms.DialogResult]::Cancel
})

$MainWindow.variables.xamlWindowControlClose.Add_Click({
    $MainWindow.window.DialogResult = [System.Windows.Forms.DialogResult]::Cancel
})

$MainWindow.variables.xamlAboutLink.Add_Click({
    [System.Diagnostics.Process]::Start($MainWindow.variables.xamlAboutLink.Header.substring(17))
})

$MainWindow.variables.xamlBrowseButton.Add_Click({
    $openFileDialog = New-Object System.Windows.Forms.OpenFileDialog
    $openFileDialog.InitialDirectory = "%USERPROFILE%\Pictures"
    $openFileDialog.Filter = "png files (*.png)|*.png"
    $openFileDialog.FilterIndex = 2
    $openFileDialog.CheckPathExists
    If ($openFileDialog.ShowDialog() -eq [System.Windows.Forms.DialogResult]::OK) {
        $MainWindow.variables.xamlBrowseTextBox.Text = $openFileDialog.FileName
        $DataTable.inputImageFullPath = $openFileDialog.FileName
        $DataTable.imageFileName = Split-Path $openFileDialog.FileName -Leaf
        $bitmapImage = New-Object System.Windows.Media.Imaging.BitmapImage
        $bitmapImage.BeginInit()
        $bitmapImage.DecodePixelWidth = 150
        $bitmapImage.UriSource = New-Object System.Uri($DataTable.inputImageFullPath)
        $bitmapImage.EndInit()
        $MainWindow.variables.xamlInputImage.Source = $bitmapImage
        Copy-Item -Path $DataTable.inputImageFullPath -Destination $DataTable.outputImageTempPath
    }
})

$functionList = @("Flipx", "Flipy", "Rotate", "Roundrobinx", "Roundrobiny", "Roundrobinrows", "Roundrobincolumns", "Pixelate", "Rgbfilter")
ForEach ($function in $functionList) {
    $MainWindow.variables.xamlTransformationComboBox.Items.Add($function)
}
$MainWindow.variables.xamlTransformationComboBox.Add_DropDownClosed({
    $DataTable.chosenTransformation = $MainWindow.variables.xamlTransformationComboBox.SelectedItem
    If ($DataTable.chosenTransformation -in @("Flipx", "Flipy", "Rotate")) {
        $MainWindow.variables.xamlPixelSlider.Visibility = "Collapsed"
        $MainWindow.variables.xamlPixelLabel.Visibility = "Collapsed"
        $MainWindow.variables.xamlColorLabel.Visibility = "Collapsed"
        $MainWindow.variables.xamlColorComboBox.Visibility = "Collapsed"
        $MainWindow.variables.xamlPixelDisplayLabel.Visibility = "Collapsed"
    } ElseIf ($DataTable.chosenTransformation -in @("Roundrobinx", "Roundrobiny", "Roundrobinrows", "Roundrobincolumns", "Pixelate")) {
        $MainWindow.variables.xamlColorLabel.Visibility = "Collapsed"
        $MainWindow.variables.xamlColorComboBox.Visibility = "Collapsed"
        $MainWindow.variables.xamlPixelSlider.Visibility = "Visible"
        $MainWindow.variables.xamlPixelLabel.Visibility = "Visible"
        $MainWindow.variables.xamlPixelDisplayLabel.Visibility = "Visible"
    } Else {
        $MainWindow.variables.xamlPixelSlider.Visibility = "Collapsed"
        $MainWindow.variables.xamlPixelLabel.Visibility = "Collapsed"
        $MainWindow.variables.xamlPixelDisplayLabel.Visibility = "Collapsed"
        $MainWindow.variables.xamlColorLabel.Visibility = "Visible"
        $MainWindow.variables.xamlColorComboBox.Visibility = "Visible"
    }
    $MainWindow.variables.xamlPreviewButton.Visibility = "Visible"
})

$MainWindow.variables.xamlPixelSlider.Maximum = 100
$MainWindow.variables.xamlPixelSlider.LargeChange = 10
$MainWindow.variables.xamlPixelSlider.SmallChange = 1
$MainWindow.variables.xamlPixelSlider.Add_ValueChanged({
    $DataTable.pixels = [Math]::round($MainWindow.variables.xamlPixelSlider.Value)
    $MainWindow.variables.xamlPixelDisplayLabel.Content = $DataTable.pixels
})

ForEach ($color in @("Red", "Green", "Blue")) {
    $MainWindow.variables.xamlColorComboBox.Items.Add($color)
}
$MainWindow.variables.xamlColorComboBox.Add_DropDownClosed({
    If ($MainWindow.variables.xamlColorComboBox.SelectedItem -match "Red") {
        $DataTable.pixels = 1
    } ElseIf ($MainWindow.variables.xamlColorComboBox.SelectedItem -match "Green") {
        $DataTable.pixels = 2
    } Else {
        $DataTable.pixels = 3
    }
})

$MainWindow.variables.xamlPreviewButton.Add_Click({
    If (Test-Path -Path $DataTable.outputImageTempPath) {
        ./ImageEdit-amd64 $(-join ("infile", "=", '"', $($DataTable.inputImageFullPath), '"')) $(-join ("outfile", "=", '"', $($DataTable.outputImageTempPath), '"')) $(("function", $DataTable.chosenTransformation) -join "=") $(("pixels", $DataTable.pixels) -join "=")
        Invoke-Item $DataTable.outputImageTempPath
    }
})

$MainWindow.variables.xamlSaveButton.Add_Click({
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

$MainWindow.window.ShowDialog() | Out-Null
