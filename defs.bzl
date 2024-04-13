# It is a good practice to wrap upstream rules in a macro.
# If you're using the same interfaces/method signatures, you can use `map_kind` to map to the upstream rule.
# Example:
# #gazelle:map_kind js_library js_library @aspect_rules_js
