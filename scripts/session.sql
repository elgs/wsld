-- session_id

SELECT USER_SESSION.*,
USER.USER_FLAG
FROM USER_SESSION INNER JOIN USER ON USER_SESSION.USER_ID=USER.ID
WHERE USER_SESSION.ID=?