alter table groups add column image_link text default 'NULL';
update groups set image_link = 'static/groups/vegetarian.jpg' where id=1;
update groups set image_link = 'static/groups/vegan.jpg' where id=2;
update groups set image_link = 'static/groups/no_sugar.jpg' where id=3;
update groups set image_link = 'static/groups/no_gluten.jpg' where id=4;
update groups set image_link = 'static/groups/keto.jpg' where id=5;
update groups set image_link = 'static/groups/no_lactose.jpg' where id=6;
