DELETE FROM users
WHERE email IN (
  'admin@opspilot.local',
  'support@opspilot.local',
  'requester@opspilot.local',
  'manager@opspilot.local'
);

DELETE FROM ticket_categories
WHERE name IN (
  'Robot Failure',
  'Access Issue',
  'Data Issue',
  'Infrastructure',
  'Change Request'
);

DELETE FROM applications
WHERE code IN ('O365', 'CORE', 'ERP');

DELETE FROM departments
WHERE code IN ('RPA', 'OPS', 'IT', 'FIN');

DELETE FROM roles
WHERE code IN ('ADMIN', 'SUPPORT', 'REQUESTER', 'MANAGER');
