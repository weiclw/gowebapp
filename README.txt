This tool helps to create web based apps for chromium based browsers on Mac.

The chromium based browsers, such as google chrome, brave browsers, uses
separate profile directories to isolate resources in such a way that multiple
instances of the browser can be launched and killed without affecting other
instances of the same browser.

On mac, the dock icons of those instances are all the same by default.
To instinguish them, icons and description texts have to be customized.
This is done by generating Mac wrappers that this program can offer.


Directory structures:

- Mac wrapper:
  - Home dir
    - Applications
      - App Name.app
        - Info.plist
        - Contents
          - MacOS
            - launcher.sh
          - Resources

- Browser profiles dir
  - Home dir
    - profiles
      - App Name


File contents:

- Info.plist
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN"
 "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>CFBundleName</key>
    <string>My Web App</string>
    <key>CFBundleDisplayName</key>
    <string>My Web App</string>
    <key>CFBundleIdentifier</key>
    <string>com.example.mywebapp</string>
    <key>CFBundleVersion</key>
    <string>1.0</string>
    <key>CFBundleExecutable</key>
    <string>launcher.sh</string>
    <key>CFBundlePackageType</key>
    <string>APPL</string>
    <key>CFBundleIconFile</key>
    <string>MyIcon</string>
</dict>
</plist>

- launcher.sh
#!/bin/bash
"/Applications/Google Chrome.app/Contents/MacOS/Google Chrome" \
  --app="https://example.com" \
  --user-data-dir="$HOME/Library/Application Support/MyWebAppProfile"
