UPDATE users
SET password_hash = '$2a$10$BfCNugPMPzHegtgQnIpjV.Nn3m0HbmxjbRPjccfP8JMY8OtimOiAq',
    updated_at = NOW()
WHERE email IN (
  'admin@opspilot.local',
  'support@opspilot.local',
  'requester@opspilot.local',
  'manager@opspilot.local'
);
