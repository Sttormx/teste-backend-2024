FROM ruby:3.0.0-alpine3.13

LABEL maintainer="Matheus M<mmacedo@icasei.com.br>"

RUN apk --update --upgrade add \
      bash \
      sqlite-dev \
      build-base \
      tzdata && \
    apk add --no-cache gcompat && \
    rm -rf /var/cache/apk/*

ENV INSTALL_PATH /ms-rails

RUN mkdir -p $INSTALL_PATH
WORKDIR $INSTALL_PATH

# Copy Gemfile and Gemfile.lock separately
COPY Gemfile Gemfile.lock ./

# Run bundle install
RUN bundle config set --local without 'development test' && \
    bundle install --jobs "$(nproc)" --no-cache

# Copy the rest of the application code
COPY . .

RUN mkdir -p $INSTALL_PATH/tmp/pids
RUN chmod +x run.sh

CMD ["./run.sh"]