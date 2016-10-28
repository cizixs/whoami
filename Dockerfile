FROM scratch
ADD whoami /whoami
EXPOSE 8778
CMD ["/whoami"]
