config:
    skip: false
Block: 
    className: {{jsformat "<<.EntityName>>_tablerow tablerow %s" "ctx.className"}}
    children:
        - Block: 
            className: tablecell field
            body: {{jsreplace "ctx.data.<<.LabelField>>"}}
        - Block:
            className: tablecell field
            body: {{jsreplace "ctx.data.UpdatedAt"}}
        - Block:
            className: tablecell field
            children:
                - Action:
                    name: update_page_<<.EntityName | lower>>
                    params: 
                        entityId: {{jsreplace "ctx.data.Id"}}
                    body: View <<.EntityName>>
