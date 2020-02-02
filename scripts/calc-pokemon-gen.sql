UPDATE pokemon SET generation =
    (CASE
        WHEN id <= 151 THEN 1
        WHEN id >= 152 AND id <= 251 THEN 2
        WHEN id >= 252 AND id <= 386 THEN 3
        WHEN id >= 387 AND id <= 493 THEN 4
        WHEN id >= 494 AND id <= 649 THEN 5
        WHEN id >= 650 AND id <= 721 THEN 6
        WHEN id >= 722 AND id <= 809 THEN 7
    END)
;