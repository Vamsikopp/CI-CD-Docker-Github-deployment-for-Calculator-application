FROM ubuntu:22.04

ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update \
&& apt-get install -y --no-install-recommends \
gnome-calculator dconf-editor dbus-x11 ca-certificates \
&& apt-get update \
&& apt-get upgrade -y \
&& apt-get remove -fy \
&& apt-get autoclean -y \
&& apt-get autoremove -y \
&& rm -rf /var/lib/apt/lists/* /var/tmp/* /tmp/*

ARG UNAME=udocker
ARG UID=1000
ARG GNAME=$UNAME
ARG GID=1000
ARG GROUPS=$GNAME

RUN groupadd -g $GID $GNAME \
&& useradd --create-home -d /home/$UNAME -g $GID -u $UID $UNAME \
&& usermod -a -G $GROUPS $UNAME
USER $UNAME
WORKDIR /home/$UNAME

ENTRYPOINT ["/usr/bin/gnome-calculator"]
