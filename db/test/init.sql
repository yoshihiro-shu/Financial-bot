INSERT INTO categories (id, name, created_at, updated_at) VALUES
(1, 'Finance', '2023-10-13 12:00:00', '2023-10-13 12:00:00'),
(2, 'Stock', '2023-10-13 12:05:00', '2023-10-13 12:05:00'),
(3, 'Currency', '2023-10-13 12:10:00', '2023-10-13 12:10:00');


INSERT INTO providers (id, name, created_at, updated_at) VALUES
(1, 'Google', '2023-10-13 12:00:00', '2023-10-13 12:00:00'),
(2, 'Yahoo', '2023-10-13 12:05:00', '2023-10-13 12:05:00'),
(3, 'Nikkei', '2023-10-13 12:10:00', '2023-10-13 12:10:00');


INSERT INTO news (id, title, description, link, thumbnail, score, published_at, provider_id, category_id) VALUES
(1, 'News Title 1', 'News Description 1', 'News Link 1', 'News Thumbnail 1', 100, '2023-10-13 12:00:00', 1, 1),
(2, 'News Title 2', 'News Description 2', 'News Link 2', 'News Thumbnail 2', 90, '2023-10-13 12:00:00', 2, 2),
(3, 'News Title 3', 'News Description 3', 'News Link 3', 'News Thumbnail 3', 80, '2023-10-13 12:00:00', 3, 3);
