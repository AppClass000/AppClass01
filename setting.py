from sqlalchemy.orm import  declarative_base, sessionmaker
from sqlalchemy import create_engine

Base = declarative_base()
engine = create_engine('mysql+mysqlconnector://root:sk3316624@localhost/appclass')
Session_local = sessionmaker(autocommit=False,autoflush=True,bind=engine)
