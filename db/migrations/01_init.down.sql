-- Drop the trigger first (if it exists)
DROP TRIGGER IF EXISTS update_products_updated_at ON products;

-- Drop the trigger function (if it exists)
DROP FUNCTION IF EXISTS update_updated_at;

-- Drop the products table
DROP TABLE IF EXISTS warehouses;
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS shops;
DROP TABLE IF EXISTS stock;
