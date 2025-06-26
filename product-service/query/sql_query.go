package query

const PRODUCT_QUERY = "SELECT p.id as Id ,p.name as product_name ,p.description as product_description ,p.available_quantity,p.price,c.id as category_id,c.name as category_name,c.description as category_description  FROM product p join category c on c.id = p.category_id "
