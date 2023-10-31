ARG PHP_VERSION=8.2
FROM php:${PHP_VERSION}-apache

ARG PHP_EXTRA_EXTENSIONS=()
ENV APACHE_DOCUMENT_ROOT=/var/www/html/public

RUN sed -ri -e 's!/var/www/html!${APACHE_DOCUMENT_ROOT}!g' /etc/apache2/sites-available/*.conf
RUN a2enmod rewrite

RUN curl -sSL https://github.com/mlocati/docker-php-extension-installer/releases/latest/download/install-php-extensions -o - | sh -s \
      pdo_mysql amqp ${PHP_EXTRA_EXTENSIONS}

RUN apt-get update && \
    apt-get install -yf unzip git && \
    rm -rf /var/lib/apt/lists/*

COPY --from=composer:2 /usr/bin/composer /usr/bin/composer