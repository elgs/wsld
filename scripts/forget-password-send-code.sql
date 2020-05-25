-- username or email

set @username := ?;

update USER SET USER.STATUS=JSON_SET(USER.STATUS, '$.rp', '$recovering-password$') 
WHERE USER.USERNAME=@username OR USER.EMAIL=@username;

SELECT EMAIL,JSON_UNQUOTE(JSON_EXTRACT(STATUS,'$.rp')) AS V_CODE  
FROM USER WHERE (USERNAME=@username OR EMAIL=@username);