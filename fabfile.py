#!/usr/bin/env python
from fabric.api import run, env, local, cd, put

env.hosts = ['45.32.24.87']
env.user = 'root'


def deploy():
    name = 'blog'
    local('GOOS=linux go build -o {0}'.format(name))
    local('gzip -f {0}'.format(name))
    with cd('/mnt/code/{0}'.format(name)):
        put('config.json', "config.json")
        put('{0}.gz'.format(name), '{0}.gz'.format(name))
        run('gzip -df {0}.gz'.format(name))
        run('chmod +x {0}'.format(name))
        run('supervisorctl restart {0}'.format(name))
    local('rm {0}.gz'.format(name))


def migrate():
    with cd('/mnt/code/blog'):
        put('migrations/blog.sql', '.')
        run('mysql -D blog < blog.sql')
