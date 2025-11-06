# MongoDB Runtime Collection Parameter

All MongoDB tools now support specifying the collection at runtime instead of requiring it in the tool configuration.

## Overview

Previously, all MongoDB tools required a `collection` field in the tool configuration, making collection selection static. Agents had to use separate tool entries for each collection.

Now, the `collection` field is **optional** in tool configuration:
- If specified in config: The tool operates on that fixed collection (backward compatible)
- If omitted from config: The `collection` becomes a required runtime parameter, allowing agents to dynamically select collections

## Affected Tools

All 9 MongoDB data manipulation tools support this feature:
- `mongodb-find`
- `mongodb-find-one`
- `mongodb-insert-one`
- `mongodb-insert-many`
- `mongodb-delete-one`
- `mongodb-delete-many`
- `mongodb-update-one`
- `mongodb-update-many`
- `mongodb-aggregate`

## Configuration Examples

### Static Collection (Existing Behavior - Still Supported)

```yaml
tools:
  query_orders:
    kind: mongodb-find
    source: my-mongodb
    database: mydb
    collection: orders  # Fixed collection
    filterPayload: '{ "customer_id": {{ .customer_id }} }'
    filterParams:
      - name: customer_id
        type: integer
    limit: 10
```

Agent invokes with:
```json
{
  "customer_id": 123
}
```

### Runtime Collection (New Feature)

```yaml
tools:
  dynamic_query:
    kind: mongodb-find
    source: my-mongodb
    database: mydb
    # collection omitted - becomes a runtime parameter
    filterPayload: '{ "customer_id": {{ .customer_id }} }'
    filterParams:
      - name: customer_id
        type: integer
    limit: 10
```

Agent invokes with:
```json
{
  "collection": "orders",
  "customer_id": 123
}
```

## Use Cases

1. **Natural-language agents**: Agent inspects user query "show me orders for customer X" and dynamically selects the `orders` collection

2. **Single tool for multiple collections**: One MCP tool can query across many collections without pre-registering separate tools

3. **Ad-hoc queries**: Agents can run exploratory queries across different collections based on context

## Implementation Details

- The `collection` parameter is automatically added as the first parameter when omitted from config
- Parameter is marked as required
- Full backward compatibility: existing configurations with static collections continue to work unchanged
- Validation ensures collection is provided either in config or at runtime

## Testing

Integration tests verify both modes:
- Static collection configuration (existing tests remain unchanged)
- Runtime collection parameter (new test added)

## Related Issue

Resolves: https://github.com/googleapis/genai-toolbox/issues/1679
