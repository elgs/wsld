-- username or email, new password, verification code
set @username := ?;
set @salt := SHA2(RAND(), 512);

UPDATE USER SET 
USER.PASSWORD=ENCRYPT(?, CONCAT('\$6\$rounds=5000$',@salt)),
USER.STATUS=JSON_REMOVE(USER.STATUS,'$.rp')
WHERE (USER.USERNAME=@username OR USER.EMAIL=@username) 
AND JSON_CONTAINS(USER.STATUS, JSON_QUOTE(?), '$.rp');