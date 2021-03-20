### 数据迁移

- 情景 1：批量迁 GB 的记录到 five_eyes

```mysql

delete from gift_order where country='india';
delete from gift_order where country='indonesia';
delete from gift_order where country='middle_east';
delete from gift_order where country='five_eyes';
delete from gift_order where country='default';
##default(IN)
insert into gift_order(`gift_id`, `order`, `scene`, `country`) SELECT gift_order.gift_id,gift_order.order,gift_order.scene, "default" FROM `gift_order` where country="IN";
##five_eyes(IN)
insert into gift_order(`gift_id`, `order`, `scene`, `country`) SELECT gift_order.gift_id,gift_order.order,gift_order.scene, "five_eyes" FROM `gift_order` where country="IN";
##india(IN)
insert into gift_order(`gift_id`, `order`, `scene`, `country`) SELECT gift_order.gift_id,gift_order.order,gift_order.scene, "india" FROM `gift_order` where country="IN";
##indonesia(ID)
insert into gift_order(`gift_id`, `order`, `scene`, `country`) SELECT gift_order.gift_id,gift_order.order,gift_order.scene, "indonesia" FROM `gift_order` where country="ID";
##middle_east(SA)
insert into gift_order(`gift_id`, `order`, `scene`, `country`) SELECT gift_order.gift_id,gift_order.order,gift_order.scene, "middle_east" FROM `gift_order` where country="SA";

##default(IN)
##middle_east(SA)
##five_eyes(IN)
##india(IN)
##indonesia(ID)

```
