do
$$
    begin
        for r in 1..100000
            loop
                INSERT INTO comments (post_id, user_id, content) VALUES (749, 400, 'Super duper OMG.. So so COOL');
            end loop;
    end;
$$;