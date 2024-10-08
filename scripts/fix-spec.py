#!/usr/bin/env python3

import yaml

SPEC_PATH = 'api/openapi.yaml'

# Load the spec file
with open(SPEC_PATH, 'r') as file:
    data = yaml.load(file, Loader=yaml.CLoader)

# Traverse schemas
if 'components' in data and 'schemas' in data['components']:
    for name, schema in data['components']['schemas'].items():
        if 'properties' in schema:
            # Remove "null" item from nullable enums
            for name, prop in schema['properties'].items():
                if 'enum' in prop and None in prop['enum']:
                    prop['enum'].remove(None)
                if 'properties' in prop and 'value' in prop['properties'] and 'enum' in prop['properties']['value'] and None in prop['properties']['value']['enum']:
                    prop['properties']['value']['enum'].remove(None)

            # Fix nullable types
            nullable_types = [
                'parent_device',
                'primary_ip',
            ]

            for ntype in nullable_types:
                if ntype in schema['properties']:
                    if 'nullable' in schema['properties'][ntype]:
                        schema['properties'][ntype]['nullable'] = True

            # Fix non-nullable types
            # See: https://github.com/OpenAPITools/openapi-generator/issues/18006
            non_nullable_types = [
                'front_image',
                'rear_image',
            ]

            for ntype in non_nullable_types:
                if ntype in schema['properties']:
                    if schema['properties'][ntype]['format'] == 'binary':
                        if 'nullable' in schema['properties'][ntype]:
                            schema['properties'][ntype].pop('nullable')

            # Fix required '*_count' not returned by Netbox API
            if 'required' in schema:
                if 'device_count' in schema['required']:
                    schema['required'].remove('device_count')

                if 'devicetype_count' in schema['required']:
                    schema['required'].remove('devicetype_count')

                if 'virtualmachine_count' in schema['required']:
                    schema['required'].remove('virtualmachine_count')

# Save the spec file
with open(SPEC_PATH, 'w') as file:
    yaml.dump(data, file, Dumper=yaml.CDumper, sort_keys=False)
