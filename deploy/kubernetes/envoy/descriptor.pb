
�
messages/user.protomessages">
User
username (	Rusername
password (	RpasswordB.Z,/microservices/auth/protoc/messages;messagesJ�
  	

  

 

 C
	
 C


  	


 

  

  


  

  

 

 


 

 bproto3
�
methods/register.protomethodsmessages/user.proto"5
RegisterRequest"
user (2.messages.UserRuser",
RegisterResponse
success (RsuccessB,Z*/microservices/auth/protoc/methods;methodsJ�
  

  

 

 A
	
 A
	
  


  



 

  	

  	

  	

  	


 




 

 

 	

 bproto3
�
messages/refresh_token.protomessages"3
RefreshToken#
refresh_token (	RrefreshTokenB.Z,/microservices/auth/protoc/messages;messagesJ�
  

  

 

 C
	
 C


  


 

  

  


  

  bproto3
�
messages/access_token.protomessages"0
AccessToken!
access_token (	RaccessTokenB.Z,/microservices/auth/protoc/messages;messagesJ�
  

  

 

 C
	
 C


  


 

  

  


  

  bproto3
�
methods/login.protomethodsmessages/refresh_token.protomessages/access_token.protomessages/user.proto"2
LoginRequest"
user (2.messages.UserRuser"�
LoginResponse;
refresh_token (2.messages.RefreshTokenRrefreshToken8
access_token (2.messages.AccessTokenRaccessTokenB,Z*/microservices/auth/protoc/methods;methodsJ�
  

  

 

 A
	
 A
	
  &
	
 %
	
 


 
 


 


  

  

  

  


 




 ,

 

 '

 *+

*



%

()bproto3
�
methods/refresh.protomethodsmessages/refresh_token.protomessages/access_token.proto"M
RefreshRequest;
refresh_token (2.messages.RefreshTokenRrefreshToken"K
RefreshResponse8
access_token (2.messages.AccessTokenRaccessTokenB,Z*/microservices/auth/protoc/methods;methodsJ�
  

  

 

 A
	
 A
	
  &
	
 %


 	 


 	

  
,

  


  
'

  
*+


 




 *

 

 %

 ()bproto3
�
methods/exit.protomethodsmessages/refresh_token.proto"J
ExitRequest;
refresh_token (2.messages.RefreshTokenRrefreshToken"&
ExitResponse
status (RstatusB,Z*/microservices/auth/protoc/methods;methodsJ�
  

  

 

 A
	
 A
	
  &


  



 

  	,

  	

  	'

  	*+


 




 

 

 	

 bproto3
�
auth_service.protoauthmethods/register.protomethods/login.protomethods/refresh.protomethods/exit.proto2�
AuthService?
Register.methods.RegisterRequest.methods.RegisterResponse6
Login.methods.LoginRequest.methods.LoginResponse<
Refresh.methods.RefreshRequest.methods.RefreshResponse3
Exit.methods.ExitRequest.methods.ExitResponseB#Z!/microservices/auth/protoc;protocJ�
  

  

 

 8
	
 8
	
   
	
 
	
 
	
	 


  


 

  M

  

  (

  3K

 D

 

 "

 -B

 J

 

 &

 1H

 A

 

  

 +?bproto3