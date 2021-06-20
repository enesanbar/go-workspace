from django.urls import path
from django.views.decorators.cache import cache_page

from courses import views

urlpatterns = [
    path('mine/', views.ManageCourseListView.as_view(), name='manage_course_list'),
    path('create/', views.CourseCreateView.as_view(), name='course_create'),
    path('<int:pk>/edit/', views.CourseUpdateView.as_view(), name='course_edit'),
    path('<int:pk>/delete/', views.CourseDeleteView.as_view(), name='course_delete'),
    path('<int:pk>/module/', views.CourseModuleUpdateView.as_view(), name='course_module_update'),
    path('module/<int:module_id>/content/<slug:model_name>/create/',
         views.ContentCreateUpdateView.as_view(),
         name='module_content_create'),
    path('module/<int:module_id>/content/<slug:model_name>/<int:id>/',
         views.ContentCreateUpdateView.as_view(),
         name='module_content_update'),
    path('content/<int:id>/delete/',
         views.ContentDeleteView.as_view(),
         name='module_content_delete'),
    # TODO: cache the templates or views that are used to display course contents to students
    path('module/<int:module_id>/',
         # cache_page(60 * 15)(views.ModuleContentListView.as_view()),
         views.ModuleContentListView.as_view(),
         name='module_content_list'),
    path('module/order/', views.ModuleOrderView.as_view(), name='module_order'),
    path('content/order/', views.ContentOrderView.as_view(), name='content_order'),
    path('', views.CourseListView.as_view(), name='course_list'),
    path('subject/<slug:subject>/', views.CourseListView.as_view(), name='course_list_subject'),
    path('enroll-course/', views.StudentEnrollCourseView.as_view(), name='student_enroll_course'),
    path('courses/', views.StudentCourseListView.as_view(), name='student_course_list'),
    path('<slug:slug>/', views.CourseDetailView.as_view(), name='course_detail'),
    path('course/<int:pk>/',
         views.StudentCourseDetailView.as_view(),
         name='student_course_detail'),
    path('course/<int:pk>/<int:module_id>/',
         views.StudentCourseDetailView.as_view(),
         name='student_course_detail_module'),

]
