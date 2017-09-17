#!/usr/local/bin/python3
from fabric.api import run, env, local, cd, put

env.hosts = ['47.52.69.7']
env.user = 'root'


def dev():
    name = 'blog-back'
    local('GOOS=linux go build')
    local('gzip blog-back')
    with cd('/mnt/code/{0}'.format(name)):
        put('{0}.gz'.format(name), '{0}.gz'.format(name))
        run('gzip -df {0}.gz'.format(name))
        run('chmod +x {0}'.format(name))
        run('supervisorctl restart {0}'.format(name))
    local('rm {0}.gz'.format(name))
