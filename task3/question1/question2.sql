drop procedure if exists getBalance;
delimiter //
create procedure getBalance (in account1_ID int, in account2_ID int)
begin
    start transaction;
    select accounts.balance into @balance from accounts where id = account1_ID;
    if @balance >= 100 then
        update accounts set balance = balance - 100 where id = account1_ID;
        update accounts set balance = balance + 100 where id = account2_ID;
        insert into transactions (from_account_id, to_account_id, amount) VALUES (account1_ID, account2_ID, 100);
        commit;
    else
        rollback;
    end if;
end //
delimiter ;
call getBalance(3, 2);
