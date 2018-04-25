docker run \
-it \
--name=blog \
--env-file env.list \
-p=8000:8080 \
--rm \
blog