-- verification code
set @vCode := ?;

-- @label:delete_flag
DELETE FROM USER_FLAG WHERE
CODE='signup'
AND VALUE=@vCode
AND EXISTS(
	SELECT 1 FROM USER_SESSION WHERE USER_FLAG.USER_ID=USER_SESSION.USER_ID AND USER_SESSION.ID='__session_id'
);