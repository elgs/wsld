-- set @now_utc := CONVERT_TZ(NOW(),'System','+0:0');

-- INSERT INTO USER_SESSION_LOG
-- SELECT ID, USER_ID, USERNAME, EMAIL, TIME_LOGIN, @now_utc, IP, STATUS, MODE
-- FROM USER_SESSION WHERE ID='__session_id';

DELETE FROM USER_SESSION WHERE ID='__session_id';