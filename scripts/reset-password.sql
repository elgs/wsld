-- username or email, new password
set @username := ?;
set @password := ?;

set @salt := SHA2(RAND(), 512);

select ID INTO @uid FROM USER WHERE (USER.USERNAME=@username OR USER.EMAIL=@username);

-- @label:reset_password
UPDATE USER SET
USER.PASSWORD=ENCRYPT(@password, CONCAT('\$6\$rounds=5000$',@salt))
WHERE USER.ID=@uid;

-- @label:delete_sessions
DELETE FROM USER_SESSION WHERE USER_SESSION.USER_ID=@uid AND USER_SESSION.ID!='__session_id';
