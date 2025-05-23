#!/bin/bash

TEMPLATE_MODULE_NAME="template"
MODULE_DIR="x"
PROTO_DIR="proto"

if [ -z "$MODULE_NAME" ]; then
  echo "Error: MODULE_NAME environment variable is not set."
  echo "Please set it like this before running the script:"
  echo '  MODULE_NAME=your_module_name ./scripts/template-module.sh'
  exit 1
fi

if [ -d "$PROTO_DIR/aether/$MODULE_NAME" ]; then
  echo "Error: The module '$MODULE_NAME''s proto files are already been created..."
  exit 1
else
	cp -r $PROTO_DIR/aether/$TEMPLATE_MODULE_NAME $PROTO_DIR/aether/$MODULE_NAME
fi

if [ -d "$MODULE_DIR/$MODULE_NAME" ]; then
  echo "Error: The module '$MODULE_NAME' already been created..."
else
	cp -r $MODULE_DIR/$TEMPLATE_MODULE_NAME $MODULE_DIR/$MODULE_NAME
fi


echo "generating the '$MODULE_NAME', and related functionalities"

