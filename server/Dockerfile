FROM hyperf/hyperf as builder

RUN wget https://github.com/composer/composer/releases/download/1.8.6/composer.phar && \
    chmod u+x composer.phar && mv composer.phar /usr/local/bin/composer && \
    composer config -g repo.packagist composer https://mirrors.aliyun.com/composer
