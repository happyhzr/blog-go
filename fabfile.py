#!/usr/local/bin/python3.6
from fabric.api import run, env, local, cd, put

env.hosts = ['47.52.69.7']
env.user = 'root'

def dev():
    local("GOOS=linux go build")
    with cd('/mnt/code/blog'):
        put('blog-back', 'blog_new')
        run('supervisorctl stop blog')
        run('mv blog_new blog')
        run('chmod +x blog')
        run('supervisorctl start blog')
    local('rm -rf blog-back')