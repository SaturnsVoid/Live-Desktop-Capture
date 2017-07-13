# Live-Desktop-Capture
A Live Desktop Capture PoC using Go and WebSockets

This program consists of 3 parts

Server

Client

index.html


The Server will await connection from the client program.

The Client program will send and base64 encoded string of the desktops screenshot to the the Server

The index.html file will use Javascript to connect to the Server using Websocks and auto refresh the image as a new one arives


Some notes:
This is a slow system, It will only work as fast as the internet connection.

The Client can be configured to compress the image, removing any color to save bandwidth and time.

This is a resource intensive system, It will use a lot of CPU and bandwidth.

This is a PoC.

You could use gZip to compress the image data more.

You could use some form of encryption.

# Packages Used

  https://github.com/AllenDang/w32
  
  https://github.com/gorilla/websocket
  
  !YOU WILL NEED A C COMPILER FOR THE W32 PACKAGE!

# Terms of Use

	* Do NOT use this on any computer you do not own, or are allowed to run this on.
	* Credits must always be given, With linksback to here.
	* You may NEVER attempt to sell this, its free and open source.
	
# Other

Go is a amazing and powerful programming language. If you already haven't, check it out; https://golang.org/

# Donations
<img src="https://blockchain.info/Resources/buttons/donate_64.png"/>
<p align="center">Please Donate To Bitcoin Address: <b>1AEbR1utjaYu3SGtBKZCLJMRR5RS7Bp7eE</b></p>
