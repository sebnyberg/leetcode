CREATE FUNCTION getUserIDs(@startDate DATE, @endDate DATE, @minAmount INT) RETURNS INT AS
BEGIN
    RETURN (
        /* Write your T-SQL query statement below. */
        select count(distinct user_id)
        from purchases
        where time_stamp between @startDate and @endDate
            and amount >= @minAmount
    );
END