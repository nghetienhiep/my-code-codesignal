{\rtf1\ansi\ansicpg1252\cocoartf2513
\cocoatextscaling0\cocoaplatform0{\fonttbl\f0\fnil\fcharset0 Menlo-Regular;}
{\colortbl;\red255\green255\blue255;\red207\green51\blue43;\red255\green255\blue255;\red149\green184\blue78;
\red191\green99\blue39;\red224\green224\blue224;\red202\green202\blue202;\red156\green151\blue223;\red202\green202\blue202;
}
{\*\expandedcolortbl;;\cssrgb\c85453\c28540\c22158;\cssrgb\c100000\c100000\c100000\c0;\cssrgb\c64619\c75971\c37836;
\cssrgb\c79882\c46893\c19734;\cssrgb\c90296\c90296\c90183;\cssrgb\c83229\c83229\c83125;\cssrgb\c67634\c66867\c89921;\cssrgb\c83137\c83137\c83137;
}
\paperw11900\paperh16840\margl1440\margr1440\vieww10800\viewh8400\viewkind0
\deftab720
\pard\pardeftab720\sl420\partightenfactor0

\f0\fs28 \cf2 \cb3 \expnd0\expndtw0\kerning0
\outl0\strokewidth0 \strokec2 func\cf4 \strokec4  \cf5 \strokec5 inversePermutation\cf6 \strokec6 (\cf4 \strokec4 permutation \cf6 \strokec6 []\cf4 \strokec4 int\cf6 \strokec6 )\cf4 \strokec4  \cf6 \strokec6 []\cf4 \strokec4 int \cf6 \strokec6 \{\cf7 \strokec7 \
\pard\pardeftab720\sl420\partightenfactor0
\cf4 \strokec4     a \cf6 \strokec6 :=\cf4 \strokec4  \cf6 \strokec6 make([]\cf4 \strokec4 int\cf6 \strokec6 ,\cf4 \strokec4  \cf6 \strokec6 len(\cf4 \strokec4 permutation\cf6 \strokec6 ))\cf7 \strokec7 \
\cf4 \strokec4     \cf2 \strokec2 for\cf4 \strokec4  i \cf6 \strokec6 :=\cf4 \strokec4  \cf2 \strokec2 range\cf4 \strokec4  permutation \cf6 \strokec6 \{\cf7 \strokec7 \
\cf4 \strokec4         pos \cf6 \strokec6 :=\cf4 \strokec4  permutation\cf6 \strokec6 [\cf4 \strokec4 i\cf6 \strokec6 ]\cf4 \strokec4  \cf6 \strokec6 -\cf4 \strokec4  \cf8 \strokec8 1\cf7 \strokec7 \
\cf4 \strokec4         a\cf6 \strokec6 [\cf4 \strokec4 pos\cf6 \strokec6 ]\cf4 \strokec4  \cf6 \strokec6 =\cf4 \strokec4  i\cf8 \strokec8 +1\cf7 \strokec7 \
\cf4 \strokec4     \cf6 \strokec6 \}\cf7 \strokec7 \
\cf4 \strokec4     \cf2 \strokec2 return\cf4 \strokec4  a\cf7 \strokec7 \
\pard\pardeftab720\sl420\partightenfactor0
\cf6 \strokec6 \}\cf9 \cb1 \strokec9 \
\
}