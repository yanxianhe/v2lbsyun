#!/usr/bin/env bash
set -ex
# mv ~/.local/lib/python3.10/site-packages/fastapi/openapi/docs.py ~/.local/lib/python3.10/site-packages/fastapi/openapi/docs_bak.py
# cp ./docs.py ~/.local/lib/python3.10/site-packages/fastapi/openapi/
uvicorn main:app --reload --host 0.0.0.0 --port 8000
