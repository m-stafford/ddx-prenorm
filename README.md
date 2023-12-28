# DDx Pre-Normalization

## Overview
One of the driving concepts behind the Democratic Data Exchange is data normalization - the process of taking incoming survey and field data and transforming it in a way that obsfucates the origin of the data and allows the data to be more easily analyzed regardless of the collecting organization. The process of normalization is tedious and involves a large amount of human hours to interpret and normalize question and response data. Incoming data cannot be ingested until this process is complete.

During my time at DDx I had always wanted a way to "pre-normalize" data so that the analysts spent their time on verification and corner cases rather than on every single incoming question. There are a lot of cute ways to do this with Natural Language Processing but with the recent general availability of Large Language Models, I wanted to see if we could apply something like ChatGPT to create a rubric, in prose, on how to handle this initial normalization.

This project is a proof-of-concept of that idea and does the following:
- Ingests a JSON file of survey questions and responses
- Uses a prompt template to query ChatGPT and requests normalization
- Transforms the response from a returned JSON object to a Golang `struct`

Normally I wouldn't bother putting a toy project like this online but I think the concept is interesting enough and hope an organization might find it useful, someday, when trying to expedite data ingestion into DDx.

## Notes
1. I've never used Golang before
2. I'm writing this on an airplane to visit family for the holidays
3. This is just a proof of concept on how we can use LLMs for data normalization and integrity

