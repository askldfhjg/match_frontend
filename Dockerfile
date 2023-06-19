FROM alpine
ADD match_frontend /match_frontend
ENTRYPOINT [ "/match_frontend" ]
