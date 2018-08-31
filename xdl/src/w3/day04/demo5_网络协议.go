package day04

/*
一.网络协议 参见: https://zh.wikipedia.org/wiki/%E7%BD%91%E7%BB%9C%E4%BC%A0%E8%BE%93%E5%8D%8F%E8%AE%AE
	1.应用层: HTTP/FTP/TFTP/SMTP/SNMP/TELNET/HTTPS/POP3/DHCP
	2.传输层: TCP/UDP
	3.网络层: ICMP/IGMP/IP/ARP/RARP
二.IPV4
	1.分类 参见: https://zh.wikipedia.org/wiki/%E5%88%86%E7%B1%BB%E7%BD%91%E7%BB%9C
		A类: 1.0.0.1		~	126.255.255.254		政府机构
		B类: 128.0.0.0	~	191.255.0.0			跨国组织
		C类: 192.0.0.0	~	223.255.255.0		商用+民用
		D类: 224.0.0.0	~	239.255.255.255		组播
		E类: 240.0.0.0	~	255.255.255.254		实验
	2.特殊地址
		1.127.x.x.x			本地网络
		2.224.x.x.x			多播地址
		3.255.255.255.255	通用广播地址
	3.Note
		1.x.x.x.0		首位IP: 当前子网网络地址
		2.x.x.x.255		末尾IP: 当前子网广播地址
三.HTTP协议
	1.HTTP请求方法
		1.Post 增
		2.Delete 删
		3.Put 改
		4.Get 查
	2.状态码
		2xx 成功			eg.200 OK
		3xx 重定向		   301 Moved Permanently
		4xx 客户端错误	   404 Not Found
		5xx 服务端错误	   500 Internal Server Error
四.网络掩码: 标识IP网络地址位和主机位
	eg. 某公司子网为10.3.134.0,现要求为产品/研发/市场三个部门划分不同网段
		产品 10.3.134.0/24
		研发 10.3.134.0/25
		市场 10.3.134.0/26
https://zh.wikipedia.org/wiki/%E5%AD%90%E7%BD%91#%E7%BD%91%E7%BB%9C%E6%8E%A9%E7%A0%81 待看
*/
