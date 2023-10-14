INSERT INTO categories (name, created_at, updated_at) VALUES
('Finance', '2023-10-13 12:00:00', '2023-10-13 12:00:00'),
('Stock', '2023-10-13 12:05:00', '2023-10-13 12:05:00'),
('Currency', '2023-10-13 12:10:00', '2023-10-13 12:10:00');


INSERT INTO providers (name, created_at, updated_at) VALUES
('Google', '2023-10-13 12:00:00', '2023-10-13 12:00:00'),
('Yahoo', '2023-10-13 12:05:00', '2023-10-13 12:05:00'),
('Nikkei', '2023-10-13 12:10:00', '2023-10-13 12:10:00');


INSERT INTO news (title, description, link, thumbnail, score, published_at, provider_id, category_id) VALUES
('News Title 1', 'News Description 1', 'News Link 1', 'News Thumbnail 1', 100, '2023-10-13 12:00:00', 1, 1),
('News Title 2', 'News Description 2', 'News Link 2', 'News Thumbnail 2', 90, '2023-10-13 12:00:00', 2, 2),
('News Title 3', 'News Description 3', 'News Link 3', 'News Thumbnail 3', 80, '2023-10-13 12:00:00', 3, 3);
