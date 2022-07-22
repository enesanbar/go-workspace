#!/usr/bin/env python
# -*- coding: UTF-8 -*-
import os
from setuptools import setup, find_packages

setup(
    name='rest',
    version='1.0',
    license='GNU General Public License v3',
    author='Enes Anbar',
    author_email='enesanbar@gmail.com',
    description='Hello world application for Flask',
    packages=find_packages(),
    platforms='any',
    install_requires=[
        'main',
        'flask',
        'Flask-Injector',
        'ccy',
    ],
    classifiers=[
        'Development Status :: 4 - Beta',
        'Environment :: Web Environment',
        'Intended Audience :: Developers',
        'License :: OSI Approved :: GNU General Public License v3',
        'Operating System :: OS Independent',
        'Programming Language :: Python',
        'Topic :: Internet :: WWW/HTTP :: Dynamic Content',
        'Topic :: Software Development :: Libraries :: Python Modules'
    ],
)
