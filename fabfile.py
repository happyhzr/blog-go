#!/usr/bin/env python
from fabric.api import run, env, local, cd, put

env.hosts = ['47.52.69.7']
env.user = 'root'


def deploy(name='blog-back'):
    local('GOOS=linux go build')
    local('gzip -f blog-back')
    with cd('/mnt/code/{0}'.format(name)):
        put('{0}.gz'.format(name), '{0}.gz'.format(name))
        run('gzip -df {0}.gz'.format(name))
        run('chmod +x {0}'.format(name))
        run('supervisorctl restart {0}'.format(name))
    local('rm {0}.gz'.format(name))


def migrate():
    with cd('/mnt/code/blog-back'):
        put('migrations/blog.sql', '.')
        run('mysql -D blog < blog.sql')
