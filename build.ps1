param(
    [string]
    [Parameter(Mandatory = $True)]
    [ValidateNotNullorEmpty()]
    [ValidatePattern('^([0-9]|[1-9][0-9]*)\.([0-9]|[1-9][0-9]*)\.([0-9]|[1-9][0-9]*)(?:-([0-9A-Za-z-]+(?:\.[0-9A-Za-z-]+)*))?(?:\+[0-9A-Za-z-]+)?$')]
    $version,

    [switch]
    $clean
)

function Clean-Artifacts
{
    Write-Host "Cleaning old build products."
    Remove-Item * -Include bittray.exe, bittray-*.zip
    if (!$?)
    {
        Write-Host -BackgroundColor red -ForegroundColor white "Failed to remove existing artifacts; see above."
        exit 1
    }
}

if ($clean -eq $True)
{
    Clean-Artifacts
    exit 0
}

if ((Get-Command "rcedit-x64.exe" -ErrorAction SilentlyContinue) -eq $null)
{
    Write-Host -ForegroundColor red "Unable to find rcedit-x64.exe in your PATH."
    exit 1
}

Write-Host "Packing bittray version $version" -ForegroundColor green

Clean-Artifacts

Write-Host "go get..."
go get

Write-Host "go vet..."
go vet ./...
if (!$?)
{
    Write-Host -BackgroundColor red -ForegroundColor white "'go vet' failed; see above."
    exit 1
}

Write-Host "go build..."
go build -ldflags -H=windowsgui bittray.go
if (!$?)
{
    Write-Host -BackgroundColor red -ForegroundColor white "'go build' failed; see above."
    exit 1
}

Write-Host "Validating artifact..."
if ( [System.IO.File]::Exists("bittray.exe"))
{
    Write-Host "Applying rsrc metadata..."
    trap
    {
        "Error adding resource metadata: $_"
    }

    rcedit-x64.exe --set-icon .\bitbucket.ico "bittray.exe"
    rcedit-x64.exe "bittray.exe" --set-version-string "ProductName" "Bittray"
    rcedit-x64.exe "bittray.exe" --set-version-string "ProductVersion" "$version"
}
else
{
    Write-Host -BackgroundColor red -ForegroundColor white "'go build' claims to have succeeded, but there is no artifact?"
    exit 1
}

$package = "bittray-$version.zip"
Write-Host "Compressing archive ($package)..."
Compress-Archive -Path .\bittray.exe -CompressionLevel Optimal -DestinationPath $package
if (!$?)
{
    Write-Host -BackgroundColor red -ForegroundColor white "Failed to create zip package."
    exit 1
}

certUtil -hashfile "$package" sha1
if (!$?)
{
    Write-Host -BackgroundColor red -ForegroundColor white "Failed to generate SHA1 integrity checksum."
    exit 1
}
else
{
    Write-Host -BackgroundColor green -ForegroundColor white "Done!"
}
