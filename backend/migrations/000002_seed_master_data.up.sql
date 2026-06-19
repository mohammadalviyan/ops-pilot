INSERT INTO roles (code, name)
VALUES
  ('ADMIN', 'Admin'),
  ('SUPPORT', 'RPA Support'),
  ('REQUESTER', 'Requester'),
  ('MANAGER', 'Manager')
ON CONFLICT (code) DO NOTHING;

INSERT INTO departments (code, name)
VALUES
  ('RPA', 'RPA Automation'),
  ('OPS', 'Operations'),
  ('IT', 'Information Technology'),
  ('FIN', 'Finance')
ON CONFLICT (code) DO NOTHING;

INSERT INTO applications (code, name, description, owner_department_id)
SELECT 'O365', 'Office 365', 'Microsoft Office 365 services', departments.id
FROM departments
WHERE departments.code = 'IT'
ON CONFLICT (code) DO NOTHING;

INSERT INTO applications (code, name, description, owner_department_id)
SELECT 'CORE', 'Core Banking', 'Core banking operational application', departments.id
FROM departments
WHERE departments.code = 'OPS'
ON CONFLICT (code) DO NOTHING;

INSERT INTO applications (code, name, description, owner_department_id)
SELECT 'ERP', 'ERP Finance', 'Finance and accounting application', departments.id
FROM departments
WHERE departments.code = 'FIN'
ON CONFLICT (code) DO NOTHING;

INSERT INTO ticket_categories (name, description)
VALUES
  ('Robot Failure', 'Robot execution failure or exception'),
  ('Access Issue', 'Login, account, permission, or access problem'),
  ('Data Issue', 'Input, output, or data quality problem'),
  ('Infrastructure', 'Server, network, firewall, or dependency issue'),
  ('Change Request', 'Change request for an existing automation')
ON CONFLICT (name) DO NOTHING;

INSERT INTO users (role_id, department_id, name, email, password_hash, status)
SELECT roles.id, departments.id, 'OpsPilot Admin', 'admin@opspilot.local', '$2a$10$7EqJtq98hPqEX7fNZaFWoOq0V0v7fF.9QG3VUPlcQJtHovHfysb6', 'ACTIVE'
FROM roles
JOIN departments ON departments.code = 'RPA'
WHERE roles.code = 'ADMIN'
ON CONFLICT (email) DO NOTHING;

INSERT INTO users (role_id, department_id, name, email, password_hash, status)
SELECT roles.id, departments.id, 'RPA Support User', 'support@opspilot.local', '$2a$10$7EqJtq98hPqEX7fNZaFWoOq0V0v7fF.9QG3VUPlcQJtHovHfysb6', 'ACTIVE'
FROM roles
JOIN departments ON departments.code = 'RPA'
WHERE roles.code = 'SUPPORT'
ON CONFLICT (email) DO NOTHING;

INSERT INTO users (role_id, department_id, name, email, password_hash, status)
SELECT roles.id, departments.id, 'Business Requester', 'requester@opspilot.local', '$2a$10$7EqJtq98hPqEX7fNZaFWoOq0V0v7fF.9QG3VUPlcQJtHovHfysb6', 'ACTIVE'
FROM roles
JOIN departments ON departments.code = 'OPS'
WHERE roles.code = 'REQUESTER'
ON CONFLICT (email) DO NOTHING;

INSERT INTO users (role_id, department_id, name, email, password_hash, status)
SELECT roles.id, departments.id, 'Ops Manager', 'manager@opspilot.local', '$2a$10$7EqJtq98hPqEX7fNZaFWoOq0V0v7fF.9QG3VUPlcQJtHovHfysb6', 'ACTIVE'
FROM roles
JOIN departments ON departments.code = 'OPS'
WHERE roles.code = 'MANAGER'
ON CONFLICT (email) DO NOTHING;
