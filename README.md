# Cmusic
A micro service for music uploading and downloading based on go micro framework.  

一个基于go-micro框架的音乐上传、下载微服务。



	Explanation:

There are three services, songList, uploadSong and downloadSong, which are registered on the consul.

The whole project includes the code implementation of the server and client of each service.

	说明：

共包括songList、uploadSong、downloadSong三个服务，将它们注册到了consul上。

整个项目包括了每个服务的服务端和客户端的代码实现。

	Installation：
git clone https://github.com/heejinzzz/Cmusic.git

	Deploy the server（部署服务端）：
以部署服务songList的服务端为例：

在安装并启动consul后，执行：

cd Cmusic/songList

go run main.go

服务 uploadSong 和 downloadSong 的服务端的部署同理。

	Client use（使用客户端）：
完成 git clone 后，在client文件夹下的client.go文件中调用songList（获取C-Music中的歌曲列表）、uploadSong（向C-Music上传歌曲）、downloadSong（从C-Music下载歌曲到本地）函数，并运行client.go即可。
