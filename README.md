# Docswave Plugin for Steampipe
## Installing and Testing the Plugin

To install the plugin, simple run the following command.

```
% make local
go build -o  ~/.steampipe/plugins/local/docswave/docswave.plugin *.go
```

Check your local plugin using the following command.

```
% steampipe plugin list
+--------------------------------------------------+---------+-------------+
| Name                                             | Version | Connections |
+--------------------------------------------------+---------+-------------+
| hub.steampipe.io/plugins/turbot/aws@latest       | 0.57.0  | aws         |
| hub.steampipe.io/plugins/turbot/steampipe@latest | 0.2.0   | steampipe   |
| local/docswave                                   | local   |             |
+--------------------------------------------------+---------+-------------+
```

Copy the sample `docswave.spc` file to `~/.steampipe/config` folder and change the name of the `plugin` from `docswave` to `local/docswave`. and update it with your token

```
% cat ~/.steampipe/config/docswave.spc
connection "docswave" {
    plugin = "local/docswave"
    
    token = "YOUR_API_TOKEN_HERE"
}
```

Check and see if you have a valid connection.

```
% steampipe plugin list
+--------------------------------------------------+---------+-------------+
| Name                                             | Version | Connections |
+--------------------------------------------------+---------+-------------+
| hub.steampipe.io/plugins/turbot/aws@latest       | 0.57.0  | aws         |
| hub.steampipe.io/plugins/turbot/steampipe@latest | 0.2.0   | steampipe   |
| local/docswave                                   | local   | docswave    |
+--------------------------------------------------+---------+-------------+
```

3 tables supported

```
+--------------------------+-------------+
| table                    | description |
+--------------------------+-------------+
| docswave_member          |             |
| docswave_team            |             |
| docswave_vacation_member |             |
+--------------------------+-------------+
```
Let's test the plugin.

```
% steampipe query "select count(member_id) from docswave_member where member_status ='WORKING'" --timing
+-------+
| count |
+-------+
| 403   |
+-------+
Time: 1.729330413s
```

That's it.

## Caution
  - If your query retrieve team_id on docswave_member then steampipe additionally calls docswave API per each member. because member API does not provide team_id.
    - e.g : If you query "select * from docswave_member", It takes over 45 sec for query 403 members.
