UPDATE users
SET password_hash = '$2a$10$7EqJtq98hPqEX7fNZaFWoOq0V0v7fF.9QG3VUPlcQJtHovHfysb6',
    updated_at = NOW()
WHERE email IN (
  'admin@opspilot.local',
  'support@opspilot.local',
  'requester@opspilot.local',
  'manager@opspilot.local'
);
