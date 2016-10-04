FROM pachyderm/job-shim:latest

# get up pip, vim, etc.
RUN apt-get -y update --fix-missing
RUN apt-get install -y python-pip python-dev libev4 libev-dev gcc libxslt-dev libxml2-dev libffi-dev vim curl
RUN pip install --upgrade pip

# get numpy, scipy, scikit-learn and flask
RUN apt-get install -y python-numpy python-scipy
RUN pip install pandas
RUN pip install scikit-learn

# add our project
ADD . /

