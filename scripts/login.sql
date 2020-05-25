-- username or email, username or email, password

set @safe_id := REPLACE(UUID(),'-','');
set @now_utc := CONVERT_TZ(NOW(),'System','+0:0');
set @username=?;

insert INTO USER_SESSION
select @safe_id, ID,USERNAME,EMAIL,@now_utc,'__client_ip',STATUS,MODE
FROM USER WHERE (USERNAME=@username OR EMAIL=@username) 
AND PASSWORD=ENCRYPT(?, PASSWORD);

SELECT * FROM USER_SESSION WHERE ID=@safe_id;