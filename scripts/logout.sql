
-- set @now_utc := CONVERT_TZ(NOW(),'System','+0:0');

-- @label:delete_session
DELETE FROM USER_SESSION WHERE ID='__session_id';
