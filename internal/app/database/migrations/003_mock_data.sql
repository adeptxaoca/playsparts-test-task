INSERT INTO manufacturers (id, name) VALUES
    (1, '4u'), (2, 'ABS'), (3, 'ACDelco'), (4, 'Acme-Trump'), (5, 'AKEBONO');

INSERT INTO parts (id, manufacturer_id, name, vendor_code) VALUES
    (1, 1, 'engine-1', 'WTLTTQJ5WW'),
    (2, 1, 'tyre-1', 'YI9CNI8VAY'),
    (3, 2, 'tyre-2', '8QQAOXJZBV'),
    (4, 3, 'windshield-3', '0ZWJ6UWF0V'),
    (5, 4, 'ressort-4', '0269JWRR59'),
    (6, 4, 'windshield-4', 'TP07V26P4P'),
    (7, 4, 'seat-4', 'DJC88XDTGX'),
    (8, 5, 'seat_5', '786V0LBRYM');
---- create above / drop below ----
DELETE FROM parts WHERE id <= 8;

DELETE FROM manufacturers WHERE id <= 5