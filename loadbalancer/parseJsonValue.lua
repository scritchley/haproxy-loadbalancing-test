core.Alert("parseJsonValue loaded");
 
function parseJsonValue(txn, salt)
     return string.match(txn.req:dup(), '"value":".*"')
end
 
core.register_fetches("parseJsonValue", parseJsonValue)
