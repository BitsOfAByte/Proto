/*
Copyright © 2022 BitsOfAByte

GPLv3 License, see the LICENSE file for more information.
*/
package main

import (
	"BitsOfAByte/proto/core"
	"io"
	"os"
)

var build_dir = "./.build_data/"

// Tasks to run as a pre-build hook
func main() {
	cleanup()
	createBuildDir(build_dir)

	generateAPTRepoFile()
	generateDNFRepoFile()
	generateDesktop()
	generateMetainfo()
	generateIcon()
}

// Create the build directory if it doesn't exist
func createBuildDir(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0755)
	}
}

// Create a file in the build directory
func createBuildFile(fileName string, data string) {
	file, err := os.Create(build_dir + fileName)
	core.CheckError(err)

	defer file.Close()

	_, err = file.WriteString(data)
	core.CheckError(err)

	file.Sync()
}

// Generate the .desktop file
func generateDesktop() {

	version := os.Args[1]

	if version == "" {
		panic("No version specified")
	}

	fileData := `[Desktop Entry]
Type=Application
Name=Proto
Comment=Proto compatability & runner manager
Icon=/usr/share/icons/proto/icon.png
Exec=proto
Categories=ConsoleOnly;Utility;X-GNOME-Utilities;FileTools;
Keywords=proton;steamplay;wine;runner;
NoDisplay=true
`

	createBuildFile("dev.bitsofabyte.proto.desktop", fileData)
}

func generateDNFRepoFile() {
	fileData := `[BitsOfAByte]            
name=BitsOfAByte Packages         
baseurl=https://packages.bitsofabyte.dev/yum/
enabled=1
gpgcheck=0`
	createBuildFile("bitsofabyte.repo", fileData)
}

func generateAPTRepoFile() {
	fileData := `deb [trusted=yes] https://packages.bitsofabyte.dev/apt/ /`
	createBuildFile("bitsofabyte.list", fileData)
}

// Generate the .metainfo.xml file
func generateMetainfo() {
	fileData := `<?xml version="1.0" encoding="UTF-8"?>
<!-- Copyright 2020 BitsOfAByte -->
<component type="desktop-application">
  <id>dev.bitsofabyte.proto</id>
  <name>Proto</name>
  <developer_name>BitsOfAByte</developer_name>
  <content_rating type="oars-1.1" />
  <launchable type="desktop-id">dev.bitsofabyte.proto.desktop</launchable>
  <metadata_license>MIT</metadata_license>
  <project_license>GPL-3.0-only</project_license>
  <summary>Manage custom runner installations</summary>
  <description>
    <p>
      Install and manage custom runners with ease in the command-line. Proto is a tool for managing custom wine runners for multiple programs without the need to manually download and extract them.
    </p>
  </description>

  <provides>
    <binary>proto</binary>
  </provides>

  <screenshots>
    <screenshot type="default">
      <caption>The Main CLI Page</caption>
      <image type="source">https://raw.githubusercontent.com/BitsOfAByte/proto/main/.assets/Screenshots/main_app_screenshot.png</image>
    </screenshot>
  </screenshots>

  <recommends>
	<display_length compare="ge">medium</display_length>
	<control>keyboard</control>
	<control>pointing</control>
	<control>console</control>
  </recommends>

  <url type="homepage">https://github.com/BitsOfAByte/proto</url>
  <url type="bugtracker">https://github.com/BitsOfAByte/proto/issues</url>
  <url type="faq">https://github.com/BitsOfAByte/proto#readme</url>
  <url type="help">https://github.com/BitsOfAByte/proto#readme</url>
</component>`

	createBuildFile("dev.bitsofabyte.proto.metainfo.xml", fileData)
}

// Fetch the icon from the assets and put it in the build directory
func generateIcon() {
	srcFile, err := os.Open("./.assets/Logos/icon.png")
	core.CheckError(err)
	defer srcFile.Close()

	destFile, err := os.Create(build_dir + "icon.png")
	core.CheckError(err)
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	core.CheckError(err)

	err = destFile.Sync()
	core.CheckError(err)
}

// Cleanup the build directory
func cleanup() {
	os.RemoveAll(build_dir)
}
