-- username, email, password
set @username := ?;
set @email := ?;
set @password := ?;

set @user_safe_id := REPLACE(UUID(),'-','');
set @user_flag_safe_id := REPLACE(UUID(),'-','');
set @salt := SHA2(RAND(), 512);
set @now_utc := CONVERT_TZ(NOW(),'System','+0:0');

-- @label:signup
INSERT INTO USER SET 
ID=@user_safe_id, 
USERNAME=@username, 
EMAIL=@email, 
PASSWORD=ENCRYPT(@password, CONCAT('\$6\$rounds=5000$',@salt)), 
MODE='',
CREATED_TIME=@now_utc;

-- @label:create_flag
INSERT INTO USER_FLAG SET
ID=@user_flag_safe_id,
USER_ID=@user_safe_id,
CODE='signup',
VALUE='__signup',
PRIVATE=1,
CREATED_TIME=@now_utc;
