-- username or email, new password

set @salt := SHA2(RAND(), 512);
set @username=?;

UPDATE USER SET
USER.PASSWORD=ENCRYPT(?, CONCAT('\$6\$rounds=5000$',@salt))
WHERE (USER.USERNAME=@username OR USER.EMAIL=@username)
AND EXISTS (SELECT 1 FROM USER_SESSION  
WHERE USER_SESSION.ID='__session_id' AND USER_SESSION.MODE='root');